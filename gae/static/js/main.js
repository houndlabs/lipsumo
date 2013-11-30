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

  var redraw = function(num, cb) {
    if (!num) { num = 4 }

    $.get("/paragraphs?num=" + num, function(resp) {


      // TODO(thomas) : can you fix this crap please? I'm trying to add the
      // class fadeOutLeft immediately after the num-paragraphs links are
      // clicked, then grab the next text, and add fadeInRight.
      $(".canvas").removeClass('fadeInRight').addClass('fadeOutLeft');
      $(".canvas").html(_.map(resp.Data, function(p) {
        return "<p>" + p + "</p>"
      }));
      $(".canvas").removeClass('fadeOutLeft').addClass('fadeInRight');

      $(".author-info .author").text(resp.Author);
      $(".author-info .title a").text(resp.Title).attr(
        'href', 'http://www.gutenberg.org/ebooks/' + resp.Id);

      if (cb) { cb() }
    });
  };

  $('.list-group-item').click(function() {
    $('.list-group-item').closest('li').removeClass('active');
    $(this).closest('li').addClass('active');
  });

  $('.repeat').click(function(e) {
    e.preventDefault();
    redraw($('li.active > a').attr('num'));
  });

})(jQuery);
