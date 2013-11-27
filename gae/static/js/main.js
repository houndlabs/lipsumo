(function($) {

  var select_text = function() {
    if (window.getSelection) {
      var range = document.createRange();
      range.selectNode($(".canvas")[0]);
      window.getSelection().addRange(range);
    }
  };

  $(".canvas").click(select_text);

  // -- FitText.js to make LIPSUMO use all the pixels

  $('.giant').fitText(0.3825);

  // -- Bootstrap helpers

  $('[rel=popover]').popover({
      html:true,
      placement:'bottom',
      content:function(){
          return $($(this).data('contentwrapper')).html();
      }
  }).on('shown.bs.popover', function() {
    var self = this;
    $(".num-paragraphs .amount a").click(function() {
      redraw($(this).attr('num'), function() {
        $(self).popover('hide');
      });
    });
  });

  $(".canvas").popover({
    html:true,
    placement: 'top',
    content: function() { return $(".book-info").html() }
  });

  // -- Page refresh

  var redraw = function(num, cb) {
    if (!num) { num = 4 }
    $.get("/paragraphs?num=" + num, function(resp) {
      $(".canvas").html(_.map(resp.Data, function(p) {
        return "<p>" + p + "</p>"
      }));

      if (cb) { cb() }
    });
  };

})(jQuery);
