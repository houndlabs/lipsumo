(function($) {

  // -- Bootstrap helpers

  $('[rel=popover]').popover({
      html:true,
      placement:'bottom',
      content:function(){
          return $($(this).data('contentwrapper')).html();
      }
  });

  // -- Page refresh

  var redraw = function(num) {
    if (!num) { num = 4 }
    $.get("/paragraphs?num=" + num, function(resp) {
      $(".canvas").html(_.map(resp.Data, function(p) {
        return "<p>" + p + "</p>"
      }));
    });
  };

  $(".num-paragraphs li a").click(function() { redraw($(this).attr('num')); });
})(jQuery);
