
$('.deposit-div1').on('click', 'li', function () {
    $(this).addClass('deposit-nuv-shadow').siblings().removeClass('deposit-nuv-shadow');
})

$('.deposit-account').on('click', 'li', function () {

    if (!$(this).hasClass('deposit-account-click')) {

        $(this).addClass('deposit-account-click').siblings().removeClass('deposit-account-click');
        $('.deposit-account').children().stop(true, true);
        var list = $('.deposit-account').children();

        var index = list.index($(this));

        list.each(function (i) {
            if (i == index) {
                $(this).animate({
                    width: "110px",
                    height: "56px",
                    top: '-=3px',
                    left: '-=5px'
                })
            } else {
                $(this).animate({
                    width: "100px",
                    height: "48px",
                    top: '17px',
                    left: i * 110 + 38 + 'px'
                })
            }
        })
    }
})
