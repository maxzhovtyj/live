// Global Variables
let localClientStream;
let webSocket;
let localClientVideo;
let videoContainer;
let peerRefs = {};

let myUID;

const muteAudioButton = document.getElementById('mute-microphone-btn');
const muteVideoButton = document.getElementById('turn-off-camera-btn');
const timeNowText = document.getElementById('time-now-text');

let time = new Date();
timeNowText.innerText = time.toLocaleString('en-US', {hour: 'numeric', minute: 'numeric', hour12: false})

muteAudioButton.addEventListener('click', () => {
    if (!localClientStream) {
        console.log('not initialized stream')
    }

    const audioTracks = localClientStream.getAudioTracks();
    if (audioTracks.length > 0) {
        const isEnabled = audioTracks[0].enabled;
        if (isEnabled) {
            muteAudioButton.innerHTML = `<svg xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 24 24" fill="#000000" class="Hdh4hc cIGbvc"><path d="M0 0h24v24H0zm0 0h24v24H0z" fill="none"></path><path d="M19 11h-1.7c0 .74-.16 1.43-.43 2.05l1.23 1.23c.56-.98.9-2.09.9-3.28zm-4.02.17c0-.06.02-.11.02-.17V5c0-1.66-1.34-3-3-3S9 3.34 9 5v.18l5.98 5.99zM4.27 3L3 4.27l6.01 6.01V11c0 1.66 1.33 3 2.99 3 .22 0 .44-.03.65-.08l1.66 1.66c-.71.33-1.5.52-2.31.52-2.76 0-5.3-2.1-5.3-5.1H5c0 3.41 2.72 6.23 6 6.72V21h2v-3.28c.91-.13 1.77-.45 2.54-.9L19.73 21 21 19.73 4.27 3z"></path></svg>`
        } else {
            muteAudioButton.innerHTML = `<svg focusable="false" width="24" height="24" viewBox="0 0 24 24" class="Hdh4hc cIGbvc NMm5M"><path d="M12 14c1.66 0 3-1.34 3-3V5c0-1.66-1.34-3-3-3S9 3.34 9 5v6c0 1.66 1.34 3 3 3z"></path><path d="M17 11c0 2.76-2.24 5-5 5s-5-2.24-5-5H5c0 3.53 2.61 6.43 6 6.92V21h2v-3.08c3.39-.49 6-3.39 6-6.92h-2z"></path></svg>`
        }
        audioTracks[0].enabled = !isEnabled;
    }

});

// Toggle video track
muteVideoButton.addEventListener('click', () => {
    if (!localClientStream) {
        console.log('not initialized stream')
    }

    const videoTracks = localClientStream.getVideoTracks();
    if (videoTracks.length > 0) {
        const isEnabled = videoTracks[0].enabled;
        if (isEnabled) {
            muteVideoButton.innerHTML = `<svg focusable="false" width="24" height="24" viewBox="0 0 24 24" class="Hdh4hc cIGbvc NMm5M"><path d="M18 10.48V6c0-1.1-.9-2-2-2H6.83l2 2H16v7.17l2 2v-1.65l4 3.98v-11l-4 3.98zM16 16L6 6 4 4 2.81 2.81 1.39 4.22l.85.85C2.09 5.35 2 5.66 2 6v12c0 1.1.9 2 2 2h12c.34 0 .65-.09.93-.24l2.85 2.85 1.41-1.41L18 18l-2-2zM4 18V6.83L15.17 18H4z"></path></svg>`
        } else {
            muteVideoButton.innerHTML = `<svg focusable="false" width="24" height="24" viewBox="0 0 24 24" class="Hdh4hc cIGbvc NMm5M"><path d="M18 10.48V6c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2v-4.48l4 3.98v-11l-4 3.98zm-2-.79V18H4V6h12v3.69z"></path></svg>`
        }
        videoTracks[0].enabled = !isEnabled;
    }
});

window.onload = () => {
    console.log('about requesting cams');
    openCamera().then((stream) => {
        localClientVideo = document.getElementById('localClientVideo');
        localClientVideo.srcObject = stream;
        localClientStream = stream;

        videoContainer = document.getElementById('video-container');
    }).then(() => {
        InitiateMeeting().then();
    });
};

const openCamera = async () => {
    if ('mediaDevices' in navigator && 'getUserMedia' in navigator.mediaDevices) {
        const allDevices = await navigator.mediaDevices.enumerateDevices();

        const cameras = allDevices.filter((device) => device.kind === 'videoinput');

        const constraints = {
            audio: {
                advanced: [{
                    echoCancellation: true,
                    noiseSuppression: true,
                    sampleRate: 48000,
                    suppressLocalAudioPlayback: true,
                },
                ]
            },
            video: {
                deviceId: cameras[0].deviceId,
                advanced: [{
                    facingMode: "right",
                }],
            },
        };

        try {
            return await navigator.mediaDevices.getUserMedia(constraints);
        } catch (error) {
            alert(error);
        }
    }
};

async function InitiateMeeting() {
    const urlParams = new URLSearchParams(window.location.search);
    const room_id = urlParams.get('id');

    if (room_id) {
        console.log('joining a meeting');
    } else {
        console.log("bs")
        return
    }

    let wsSchema = "wss"

    if (location.href.slice(0, 5) !== 'https') {
        wsSchema = "ws"
    }

    let socket = new WebSocket(
        `${wsSchema}://${document.location.host}/ws/join-room?roomID=${room_id}`
    );

    webSocket = socket;

    myUID = document.getElementById("current-user-id").value

    socket.addEventListener('open', () => {
        socket.send(JSON.stringify({join: true, from: myUID}));
    });

    socket.addEventListener('message', async (e) => {
        const message = JSON.parse(e.data);

        console.log(message)

        let fromUserID = message.from;

        if (message.join) {
            console.log('someone just joined the call', message);
            callUser(fromUserID);
        }

        if (message.to && message.to !== myUID) {
            return
        }

        if (message.iceCandidate) {
            console.log('receiving and adding ICE candidate');
            try {
                await peerRefs[fromUserID].addIceCandidate(message.iceCandidate);
            } catch (error) {
                return
            }
        }

        if (message.offer) {
            await handleOffer(message.offer, socket, fromUserID);
        }

        if (message.answer) {
            handleAnswer(message.answer, fromUserID);
        }
    });
}

const handleOffer = async (offer, socket, remoteUID) => {
    console.log('received an offer, creating an answer');

    let peerRef = createPeer(remoteUID);

    await peerRef.setRemoteDescription(new RTCSessionDescription(offer));

    localClientStream.getTracks().forEach((track) => {
        peerRef.addTrack(track, localClientStream);
    });

    const answer = await peerRef.createAnswer();
    await peerRef.setLocalDescription(answer);

    socket.send(JSON.stringify({answer: peerRef.localDescription, from: myUID, to: remoteUID}));

    peerRefs[remoteUID] = peerRef
};

const handleAnswer = (answer, fromUID) => {
    console.log('received an answer, creating RTC session');

    peerRefs[fromUID].setRemoteDescription(new RTCSessionDescription(answer));
};

const callUser = (remoteUID) => {
    console.log('calling other remote user');
    let peerRef = createPeer(remoteUID);

    localClientStream.getTracks().forEach((track) => {
        peerRef.addTrack(track, localClientStream);
    });

    peerRefs[remoteUID] = peerRef
};

const createPeer = (remoteUID) => {
    console.log('creating peer connection');

    const peer = new RTCPeerConnection({
        iceServers: [{urls: 'stun:stun.l.google.com:19302'}],
    });

    peer.onnegotiationneeded = () => handleNegotiationNeeded(remoteUID);
    peer.onicecandidate = (e) => handleIceCandidate(e, remoteUID);
    peer.ontrack = (e) => handleTrackEvent(e, remoteUID);

    return peer;
};

const handleNegotiationNeeded = async (remoteUID) => {
    console.log('creating offer');

    try {
        const myOffer = await peerRefs[remoteUID].createOffer();
        await peerRefs[remoteUID].setLocalDescription(myOffer);
        webSocket.send(JSON.stringify({offer: peerRefs[remoteUID].localDescription, from: myUID, to: remoteUID}));
    } catch (error) {
        alert(error);
    }
};

const handleIceCandidate = (e, remoteUID) => {
    console.log('found ice candidate');
    if (e.candidate) {
        webSocket.send(JSON.stringify({iceCandidate: e.candidate, to: remoteUID, from: myUID}));
    }
};

const handleTrackEvent = (e, remoteUID) => {
    console.log('received tracks');

    if (!remoteUID) {
        console.log("tracks, remoteUID", remoteUID)
        return
    }

    let videoElement = document.getElementById(`remoteClientVideo-${remoteUID}`)

    if (!videoElement) {
        const newDiv = document.createElement('div');
        newDiv.className = 'p-2 max-w-lg';

        videoElement = document.createElement('video');
        videoElement.playsInline = true;
        videoElement.className = 'flex rounded-lg w-full h-full';
        videoElement.style.transform = 'scaleX(-1)';
        videoElement.autoplay = true;
        videoElement.id = `remoteClientVideo-${remoteUID}`

        newDiv.appendChild(videoElement);

        videoContainer.appendChild(newDiv);
    }

    videoElement.srcObject = e.streams[0];
};
