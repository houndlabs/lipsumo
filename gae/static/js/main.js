(function($) {

  $.get("/paragraphs", function(resp) {
    $(".canvas").html(_.map(resp.Data, function(p) {
      return "<p>" + p + "</p>"
    }));
  });
})(jQuery);
