$(document).ready(function(){
	$('.ui.left.sidebar').sidebar({
	    context: $('.ui.pushable.segment'),
	    transition: 'push',
	    closable: false,
	}).sidebar('attach events', '#mobile_item');

	$('.ui.right.sidebar').sidebar({
	    context: $('.ui.pushable.segment'),
	    transition: 'overlay'
	}).sidebar('attach events', '#user_item');

	$('.example .menu .browse').popup({
	  	popup : $('.mega-menu')
	  });

	$('.message .close')
	  .on('click', function() {
	    $(this)
	      .closest('.message')
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

	$('.ui.sticky')
		.sticky({
			context: '#context'
		})
	;


	function jqUpdateSize(){
	    var width = $(window).width();
	    autoHideSidebar(width);
	};

	function autoHideSidebar(width){
	    console.log(width);

		var sidebar = $('.ui.left.sidebar');
		    if (width <= 768) {
				sidebar.removeClass('visible');
		    }else{
				sidebar.addClass('visible');
				$(".pusher").removeClass("dimmed");   	
		    }
	}

	autoHideSidebar($(window).width());
	$(window).resize(jqUpdateSize);

});
