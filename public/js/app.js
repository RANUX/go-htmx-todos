
document.onreadystatechange = function () {
  var state = document.readyState

  if (state == 'interactive') {
       document.getElementById('main').style.visibility="hidden";
       let loadEl = document.getElementById('load')
       loadEl.style.visibility="visible";

       // background:url("/public/img/spinner.gif") no-repeat center center;
       loadEl.style.background="url(/public/img/spinner.gif) no-repeat center center";
       // background size 100x100
       loadEl.style.backgroundSize="50px 50px";
       // width 100%
       loadEl.style.width="100%";
       // height 100%
       loadEl.style.height="100%";
       // position fixed
       loadEl.style.position="fixed";
       // top 0
       loadEl.style.top="0";
       // left 0
       loadEl.style.left="0";
       // z-index 9999
       loadEl.style.zIndex="9999";
  } else if (state == 'complete') {
      setTimeout(function(){
         document.getElementById('interactive');
         document.getElementById('load').style.visibility="hidden";
         document.getElementById('main').style.visibility="visible";
      },1000);
  }
}