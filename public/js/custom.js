$(document).ready(function () {
	// $('.ui.left.sidebar').sidebar({
	// 	context: $('.pusher'),
	// 	transition: 'overlay',
	// 	closable: false,
	// }).sidebar('attach events', '#mobile_item');

	$('.ui.right.sidebar').sidebar({
		context: $('.pusher'),
		transition: 'overlay'
	}).sidebar('attach events', '#user_item');

	$('.example .menu .browse').popup({
		popup: $('.mega-menu')
	});

	$('.message .close')
		.on('click', function () {
			$(this)
				.closest('.message')
				.parent()
				.transition('fade')
				.remove()
				;
		})
		;

	$('.ui.channels.dropdown')
		.dropdown({
			allowAdditions: true,
			maxSelections: 3,
			apiSettings: {
				// this url parses query server side and returns filtered results
				url: '/c/json/?s={query}'
			},
		})
		;

	$('.ui.post-config').dropdown();

	$('.ui.sticky').sticky({
		context: '#content',
		offset: 25
	});

	function jqUpdateSize() {
		var width = $(window).width();
		autoHideSidebar(width);
	};

	function autoHideSidebar(width) {
		var sidebar = $('.ui.left.sidebar');
		if (width <= 768) {
			sidebar.removeClass('visible');
		} else {
			sidebar.addClass('visible');
			$(".pusher").removeClass("dimmed");
		}
	}

	autoHideSidebar($(window).width());
	$(window).resize(jqUpdateSize);

	$('.topic-pages').dropdown({
		onChange: function (val) {
			window.location.href = val;
		}
	});

	$('.popup').popup();

	//
	// SETTINGS
	//

	$('.header-channels.dropdown')
		.dropdown({
			allowAdditions: false,
			maxSelections: 10,
			clearable: true,
			apiSettings: {
				// this url parses query server side and returns filtered results
				url: '/c/json/?s={query}'
			},
		})
		;

	$('.todays-posts-channels.dropdown')
		.dropdown({
			allowAdditions: false,
			maxSelections: 5,
			clearable: true,
			apiSettings: {
				// this url parses query server side and returns filtered results
				url: '/c/json/?s={query}'
			},
		})
		;

	$('.customized-todays-select').checkbox({
		onChecked: function () {
			$(".customized-todays-posts").removeClass("transition hidden");
		},
		onUnchecked: function () {
			$('.todays-posts-channels.dropdown').dropdown("clear");
			$(".customized-todays-posts").addClass("transition hidden");
		}
	}
	);

});
