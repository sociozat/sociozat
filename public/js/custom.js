$(document).ready(function () {
	const STICKY_OFFSET = 160;
	document.addEventListener("htmx:after-swap", (event) => {
		const content = document.getElementById("content");
		window.scrollTo(0, content);
	});

	// $('.ui.left.sidebar').sidebar({
	// 	context: $('.pusher'),
	// 	transition: 'overlay',
	// 	closable: false,
	// }).sidebar('attach events', '#mobile_item');

	$('.ui.mbl.sidebar').sidebar({
		context: $('.pusher'),
		transition: 'push',
		direction: 'right',
		dimPage: false,
	})
		.sidebar('attach events', '#mobile_item')
		.sidebar('setting', 'transition', 'push');


	$('.ui.user.sidebar').sidebar({
		context: $('.pusher'),
		transition: 'push',
		direction: 'left',
		dimPage: false,
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

	$('.trending-posts-channels.dropdown')
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

	$('.customized-trending-select').checkbox({
		onChecked: function () {
			$(".customized-trending-posts").removeClass("transition hidden");
		},
		onUnchecked: function () {
			$('.trending-posts-channels.dropdown').dropdown("clear");
			$(".customized-trending-posts").addClass("transition hidden");
		}
	}
	);


	// SEARCH
	$('.ui.search').search({
		type: 'category',
		minCharacters: 3,
		cache: true,
		apiSettings: {
			onResponse: function (apiResponse) {
				const response = {
						results: {},
						actions: {}
					}
				;
				// translate GitHub API response to work with search
				$.each(apiResponse.results, function (index, item) {
					var maxResults = 12;
					if (index >= maxResults) {
						return false;
					}
					let type = item.type;

					// create new language category
					if (response.results[type] === undefined) {
						response.results[type] = {
							name: type,
							results: []
						};
					}

					let slug;
					if (type === "post") {
						slug = "/p/" + item.slug
					} else if (type === "topic") {
						slug = "/t/" + item.slug
					} else if (type === "channel") {
						slug = "/c/" + item.slug
					} else if (type === "user") {
						slug = "/u/" + item.slug
					}

					// add result to category
					response.results[type].results.push({
						title: item.title,
						url: slug
					});
				});
				response.actions = apiResponse.actions
				return response;
			},
			url: '/api/search/{query}'
		},
	});


	//textarea buttons
	//TODO improve this feature
	$(".editor-button").on("click", function(){
	   var text = $(this).data("content")
	   var $txt = $("textarea#content");
	   var caretPos = $txt[0].selectionStart;
	   var textAreaTxt = $txt.val();
	   $txt.val(textAreaTxt.substring(0, caretPos) + text + textAreaTxt.substring(caretPos) );
	});


	//auto scroll to body
	$(document).click(function(e) {
		$("body").scrollTop(0);
	});

});