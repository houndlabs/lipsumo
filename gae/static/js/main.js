(function($) {

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
