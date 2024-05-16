$(document).ready(function () {

	// Create a socket
	let socket = new WebSocket('wss://'+window.location.host+'/ws/posts')

	// Display a message
	let display = function(event) {
		let cids = event.Text.split(',')
		cids.push("0") //dont miss the today
		cids.forEach(function(id){
			const elem = $('.cid'+id)
			if(Object.keys(elem).length > 0){
				let total;
				let badge = elem.children().first();
				if(badge.hasClass('p-count-content')){
					total = parseInt(badge.html()) + 1;
				}else{
					badge  = $('<div/>', {
						"class": 'ui label tiny p-count-content',
					}).appendTo(elem);
					total = 1;
				}
				badge.html(total)
			}
		})
	}

	// Message received on the socket
	socket.onmessage = function(event) {
		display(JSON.parse(event.data))
	}

	$('.p-count').click(function (){
		let badge = $(this).children().first();
		if(badge.hasClass('p-count-content')){
			badge.remove()
		}
	});

	const STICKY_OFFSET = 160;
	document.addEventListener("htmx:after-swap", (event) => {
		const content = document.getElementById("content");
		window.scrollTo(0, content + STICKY_OFFSET);
	});


	$('.ui.mbl.sidebar').sidebar({
		context: $('.pusher'),
		direction: 'left',
		transition: 'push',
		silent: true,
		dimPage: false,
	})
		.sidebar('attach events', '#mobile_item')
		.sidebar('setting', 'transition', 'push');


	$('.ui.user.sidebar').sidebar({
		transition: 'push',
		silent: true,
		direction: 'right',
		dimPage: false,
	})
		.sidebar('attach events', '#user_item')
		.sidebar('setting', 'transition', 'push');


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

	htmx.onLoad(function(elt) {

		autoHideSidebar($(window).width());
		$(window).resize(jqUpdateSize);

		//
		// SETTINGS
		//
		$('.popup').popup();

		// $('.ui.sticky').sticky({
		// 	context: '.eight',
		// 	offset: 25
		// });

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


		$('.topic-pages').dropdown({
			onChange: function (val) {
				return htmx.ajax('GET', val, {target:'#content', url: true})
			}
		});

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

		$('.post-action').off().on("click",function(e){
			e.preventDefault();
			const post = $(this).attr('data-post');
			const action = $(this).attr('data-action');
			$.post( "/a/action/" + post, {"action" : action}, function (response){

				$("a[data-post='" + post + "']").each(function(){
					$this = $(this)
					const icon = $this.children().first();
					if(!icon.hasClass('outline')) icon.addClass('outline')
					const value = icon.hasClass('up') ? response.total.Likes : response.total.Dislikes;

					if(icon.hasClass('up') && action == "like"){
						icon.removeClass('outline')
					}

					if(icon.hasClass('down') && action == "dislike"){
						icon.removeClass('outline')
						console.log(icon)
					}

					$this.html(icon.prop('outerHTML') + value)
				})

			});
		});

		//textarea buttons
		//TODO improve this feature
		$(".editor-button").off().on("click", function(){
			var text = $(this).data("content")
			var $txt = $("textarea#content");
			var caretPos = $txt[0].selectionStart;
			var textAreaTxt = $txt.val();
			$txt.val(textAreaTxt.substring(0, caretPos) + text + textAreaTxt.substring(caretPos) );
		});
	})
});