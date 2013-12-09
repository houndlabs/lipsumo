(function($) {

  // -- Select all the text on click.

  var select_text = function() {
    if (!window.getSelection) { return }

    var range = document.createRange();
    range.selectNode($(".canvas")[0]);
    window.getSelection().addRange(range);
  };

  $(".canvas").click(select_text);

  // -- FitText.js to make LIPSUMO use all the pixels

  $('.giant').fitText(0.3825);

  // -- Refresh the text

  var redraw = function(num) {
    var canvas = $(".canvas");
    var metadata = $(".metadata");

    if (!num) { num = 4 }

    var update_page = function(resp) {
      canvas.html(_.map(resp.Data, function(p) {
        return "<p>" + p + "</p>"
      }));
      metadata.html(_.template($("#tmpl-metadata").html(), resp));

      canvas.removeClass('fadeOutLeft').addClass('fadeInRight');
      metadata.removeClass('fadeOutLeft').addClass('fadeInRight');
    };

    canvas.addClass('fadeOutLeft');
    metadata.addClass('fadeOutLeft');
    $.get("/paragraphs?num=" + num, update_page);
  };

  $('.list-group-item').click(function() {
    $('.list-group-item').closest('li').removeClass('active');
    $(this).closest('li').addClass('active');
    redraw($('li.active > a').attr('num'));
  });

  $('.repeat').click(function(e) {
    e.preventDefault();
    redraw($('li.active > a').attr('num'));
  });

  // -- Event tracking

  twttr.ready(function (twttr) {
    twttr.events.bind('follow', function(e) {
      analytics.track('twitter.follow', {
        category: 'twitter',
        action: 'follow'
      });

      ga('send', 'social', 'twitter', 'follow');
    });

    twttr.events.bind('tweet', function (e) {
      analytics.track('twitter.tweet', {
        category: 'twitter',
        action: 'tweet'
      });

      ga('send', 'social', 'twitter', 'tweet');
    });
  });

  $("#facebook-jssdk")[0].onload = function() {
    FB.Event.subscribe('edge.create', function() {
      analytics.track('facebook.like', {
        category: 'facebook',
        action: 'like'
      });

      ga('send', 'social', 'facebook', 'like');
    });
  };

  $(".donate > a").click(function() {
    analytics.track('charity.donate', {
      category: 'charity',
      action: 'donate',
      label: $(this).attr('data-dest')
    });
  });

  $("#mc-embedded-subscribe-form").submit(function() {
    analytics.track('email.subscribe', {
      category: 'email',
      action: 'subscribe'
    });
  });

})(jQuery);
