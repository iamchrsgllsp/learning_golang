function test() {
    alert("Hello World!");
    
}

function showMenu2(value){
  let id = document.getElementById(value);
  if (id.style.display == "none"){id.style.display = "block";} else {id.style.display = "none";}
}