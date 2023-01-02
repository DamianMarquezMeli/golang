window.onload = function() {

	var data;

	var json = (function () {
		var json = null;

		$.ajax({
			async: false,
			global: false,
			type: "POST",
			//url: "arr2.php",
			url: "php/bd_a_json.php",
			dataType: "json",
			success: function (data) {
				json = data;
				alert("existo");
			},
			error: function (data) {
				alert("ERROR");
			}
		});

		return json;
	});

	alert(json.length);
}
