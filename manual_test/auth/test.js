var v1 = 'http://localhost:3100/auth/salesforce?salesforceNamespace=c&instanceType=SANDBOX';
var v2_sf = 'http://localhost:3100/auth/v2/salesforce?salesforceNamespace=c&state=6380672431a8cdf563ff521225e07e91a83ab217eeb633e6b5acdac4499a4e822abbee4d2351cfb3c1fb721515143d9ba5df21473a80df145a7401ce760f25e567dddc1658c8ab8e67d6ad2a64994e571a315885a2168ef74a13d14fcf03fced18c0acfff95b6dbf1ffec14f3d719b625bd7cb232184c80c73f019a1243ffe29ef3e183c074372e773d6b8e4ef7f883737d3f24174d4b4a211df45e8ecb9d034';
var v2_st = 'http://localhost:3100/auth/v2/stripetestmode?salesforceNamespace=c&state=bb6363f0ebf9af9caafa7f20b33046f97f45589495e226e034fb439ac06233b3a9d0e262d554f771de292c6edcc3dd8a473adf91930373e75c82de945520256747ab232430d8a7fb128dfa762951a61deb397e592917afa559bcddd358830e3a1f866796bfec105ae9392cf75013a7819e8b3b3e044e876d9ffd6fca75d101245d6302364283a13fc304662b5a843a10d048b4ef2eb0e241b01cc5f141290a72ea529eb3c04d07c953e9f84ca99f18a193bef0c93012d72d00d8331600f3416fe06440d82b384622b27bb2db56df0b066b9f71a9879b8be93751e1fcd15cf86e';

var subwin = null;

function showV1() {
    subwin = window.open(v1, '_blank');
    return false;
}

function showV2Sf() {
    subwin = window.open(v2_sf, '_blank');
    return false;
}

function showV2St() {
    subwin = window.open(v2_st, '_blank');
    return false;
}

window.addEventListener('onpostmessage', function(event) {
    var doc = document.getElementById('postmsg');
    doc.innerHTML += '<div>' + JSON.stringify(e) + '</div>';
});