//V1版
const mousewheel = function mymousewheel(getDom) {

    getDom.addEventListener('wheel', function (e) {
        var e = window.event || e;
        mousewheelProcess(getDom, e)
    })

    function mousewheelProcess(getDom, e) {

        //滚出去的距离
        var scrollTop = getDom.scrollTop,
            //元素的总高度
            scrollHeight = getDom.scrollHeight,
            //元素的可视高度
            height = getDom.clientHeight;

        var outHeight = scrollHeight - height;
        // console.log(outHeight)

        var distanceY = scrollHeight - (scrollTop + height);
        //console.log(distanceY)


        if (e.deltaY > 0) {
            //console.log('下');
            if (distanceY <= 5) {
                if (e.preventDefault) {
                    e.preventDefault();

                }
                else {
                    e.returnValue = false;
                }
            }

        } else {

            //console.log('上');
            if (outHeight == distanceY) {
                if (e.preventDefault) {
                    e.preventDefault();

                }
                else {
                    e.returnValue = false;
                }
            }
        }
    }
}
export default mousewheel;