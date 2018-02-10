function clock(time){
	var oldtime = new Date(time);
	var severtime  =  Date.parse(oldtime)/1000;//时间戳
	var timestamp = Date.parse(new Date()); //客户端时间戳 
	timestamp = timestamp / 1000;
	var between = severtime - timestamp;//客户端与服务端时间差
	var gmt_time = timestamp + between;
	// Cache some selectors
	var clock = $('#clock'),
		alarm = clock.find('.alarm'),
		ampm = clock.find('.ampm');

	// Map digits to their names (this will be an array)
	var digit_to_name = 'zero one two three four five six seven eight nine'.split(' ');

	// This object will hold the digit elements
	var digits = {};

	// Positions for the hours, minutes, and seconds
	var positions = [
		'h1', 'h2', ':', 'm1', 'm2', ':', 's1', 's2'
	];

	// Generate the digits with the needed markup,
	// and add them to the clock

	var digit_holder = clock.find('.digits');


	$.each(positions, function(){

		if(this == ':'){
			digit_holder.append('<div class="dots">');
		}
		else{

			var pos = $('<div>');

			for(var i=1; i<8; i++){
				pos.append('<span class="d' + i + '">');
			}

			// Set the digits as key:value pairs in the digits object
			digits[this] = pos;

			// Add the digit elements to the page
			digit_holder.append(pos);
		}

	});

	// Add the weekday names

	// var weekday_names = 'MON TUE WED THU FRI SAT SUN'.split(' '),
		// weekday_holder = clock.find('.weekdays');
	// $.each(weekday_names, function(){
		// weekday_holder.append('<span>' + this + '</span>');
	// });

	// var weekdays = clock.find('.weekdays span');

		week_holder = clock.find('.display_week');
		date_holder = clock.find('.display_date');

	// Run a timer every second and update the clock
	(function update_time(){
		// Use moment.js to output the current time as a string
		// hh is for the hours in 12-hour format,
		// mm - minutes, ss-seconds (all with leading zeroes),
		// d is for day of week and A is for AM/PM
		// alert(newtime);
		time = new Date(gmt_time * 1000);
		date = time.getFullYear() + "年" + (time.getMonth()+1) + "月" + time.getDate() + "日";
		week = time.getDay();
		switch(week){
                case 1:
                    week="周一";
                    break;
                case 2:
                    week="周二";
                    break;
                case 3:
                    week="周三";
                    break;
                case 4:
                    week="周四";
                    break;
                case 5:
                    week="周五";
                    break;
                case 6:
                    week="周六";
                    break;
                case 0:
                    week="周天";
                    break;
                default:
                    week="什么情况!";
            }
		// var now = moment().format("hhmmssdA");
		week = '北京时间'+ ' ' + week;
		week_holder.text(week);
		date_holder.text(date);
	    now = moment(time).format("HHmmssdA");
		digits.h1.attr('class', digit_to_name[now[0]]);
		digits.h2.attr('class', digit_to_name[now[1]]);
		digits.m1.attr('class', digit_to_name[now[2]]);
		digits.m2.attr('class', digit_to_name[now[3]]);
		digits.s1.attr('class', digit_to_name[now[4]]);
		digits.s2.attr('class', digit_to_name[now[5]]);

		// The library returns Sunday as the first day of the week.
		// Stupid, I know. Lets shift all the days one position down, 
		// and make Sunday last

		var dow = now[6];
		dow--;
		
		// Sunday!
		if(dow < 0){
			// Make it last
			dow = 6;
		}

		// Mark the active day of the week
		// weekdays.removeClass('active').eq(dow).addClass('active');

		// Set the am/pm text:
		// ampm.text(now[7]+now[8]);
		// console.dir(date_html);

	    timestamp = Date.parse(new Date()); //客户端时间戳 
		timestamp = timestamp / 1000;
		gmt_time = timestamp + between;
		// Schedule this function to be run again in 1 sec
		setTimeout(update_time, 1000);

	})();

	// Switch the theme

	// $('a.button').click(function(){
	// 	clock.toggleClass('light dark');
	// });

}