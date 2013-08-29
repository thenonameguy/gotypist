var w = window,
	d = document;

function newEvent(element, what, callback){
	if (element.attachEvent){
		element.attachEvent('on' + what, callback);
	}
	else {
		element.addEventListener(what, callback);
	}
}

newEvent(w, "load", init_scripts);

function init_scripts(){
	var logo = d.getElementById("logo");
	
	var x = 0;
	function animate_logo(){
		x = x <= 1.0 ? x + 0.01 : 0.0;
		logo.style.opacity = Math.cos(x*Math.PI*2.0)/4 + 0.75;
	}
	var t1 = setInterval(animate_logo, 1000/60);
	
	var run_it = true;
	function toggle_link(){
		if (run_it){
			var y = 0;
			var t2 = setInterval(move_logo_up, 1000/60);
			run_it = false;
		}
		else {
			logo.onmouseover = null;
		}
		function move_logo_up(){
			if (y >= 1.0){
				clearInterval(t1);
				clearInterval(t2);
				logo.style.opacity = 1;
				toggle_link_animation();
				return;
			}
			logo.style.marginTop = ((-58-80*Math.sin(y*Math.PI/2))|0) + "px";
			y += 0.02;
		}
	}
	newEvent(logo, "mouseover", toggle_link);
	
	function toggle_link_animation(){
		var link = d.getElementById("link_holder").getElementsByTagName("a")[0];
		link.style.visibility = "visible";
		var x = 0;
		function animate_link(){
			if (x <= 1.0){
				x += 0.1;
				link.style.opacity = x;
			}
			else {
				clearInterval(t3);
			}
		}
		var t3 = setInterval(animate_link, 1000/60);
	}
}