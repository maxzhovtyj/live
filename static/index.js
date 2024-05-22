// Global Variables
let localClientStream;
let webSocket;
let localClientVideo;
let remoteClientVideo;
let peerRef;

window.onload = () => {
    console.log('about requesting cams');
    openCamera().then((stream) => {
        localClientVideo = document.getElementById('localClientVideo');
        localClientVideo.srcObject = stream;
        localClientStream = stream;

        remoteClientVideo = document.getElementById('remoteClientVideo');
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
            console.log(error);
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

    let socket = new WebSocket(
        `wss://${document.location.host}/ws/join-room?roomID=${room_id}`
    );

    webSocket = socket;

    socket.addEventListener('open', () => {
        socket.send(JSON.stringify({join: true}));
    });

    socket.addEventListener('message', async (e) => {
        const message = JSON.parse(e.data);

        if (message.join) {
            console.log('Someone just joined the call');
            callUser();
        }

        if (message.iceCandidate) {
            console.log('recieving and adding ICE candidate');
            try {
                await peerRef.addIceCandidate(message.iceCandidate);
            } catch (error) {
                console.log(error);
            }
        }

        if (message.offer) {
            await handleOffer(message.offer, socket);
        }

        if (message.answer) {
            handleAnswer(message.answer);
        }
    });
}

const handleOffer = async (offer, socket) => {
    console.log('recieved an offer, creating an answer');

    peerRef = createPeer();

    await peerRef.setRemoteDescription(new RTCSessionDescription(offer));

    localClientStream.getTracks().forEach((track) => {
        peerRef.addTrack(track, localClientStream);
    });

    const answer = await peerRef.createAnswer();
    await peerRef.setLocalDescription(answer);

    socket.send(JSON.stringify({answer: peerRef.localDescription}));
};

const handleAnswer = (answer) => {
    peerRef.setRemoteDescription(new RTCSessionDescription(answer));
};

const callUser = () => {
    console.log('calling other remote user');
    peerRef = createPeer();

    localClientStream.getTracks().forEach((track) => {
        peerRef.addTrack(track, localClientStream);
    });
};

const createPeer = () => {
    console.log('creating peer connection');
    const peer = new RTCPeerConnection({
        iceServers: [{urls: 'stun:stun.l.google.com:19302'}],
    });

    peer.onnegotiationneeded = handleNegotiationNeeded;
    peer.onicecandidate = handleIceCandidate;
    peer.ontrack = handleTrackEvent;

    return peer;
};

const handleNegotiationNeeded = async () => {
    console.log('creating offer');

    try {
        const myOffer = await peerRef.createOffer();
        await peerRef.setLocalDescription(myOffer);
        webSocket.send(JSON.stringify({offer: peerRef.localDescription}));
    } catch (error) {
        console.log(error);
    }
};

const handleIceCandidate = (e) => {
    console.log('found ice candidate');
    if (e.candidate) {
        webSocket.send(JSON.stringify({iceCandidate: e.candidate}));
    }
};

const handleTrackEvent = (e) => {
    console.log('Recieved tacks');
    remoteClientVideo.srcObject = e.streams[0];
};
