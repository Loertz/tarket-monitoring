$("#log-screen").click(function openNav() {
    $("#mySidenav").width("100%");
    /*    document.getElementById("main").style.marginLeft = "100%";*/
});

/* Set the width of the side navigation to 0 and the left margin of the page content to 0 */
$("#acti-screen").click(function closeNav() {
    $("#mySidenav").width("20%");
});

function changeImage(id, a) {
    document.getElementById(id).src = a;
};

function changeText(id,t) {
    document.getElementById(id).innerText = t;
};

function evenement (a){
    var event = document.createElement("p");
    event.innerText = a ;
    $("#eventbar").prepend(event).fadeIn() ;
};

function* idMaker(a) {
    var index = 0;
    while(true)
        yield index++;
}


function initiate(data) {
    data.forEach(function create(val) {

        var box = document.createElement("div");
        box.setAttribute("class", "floating-box");
        box.id = "room-" + val.n.toString();
        box.setAttribute('data', val.tmc);

        var info = document.createElement("div");
        info.setAttribute("class", "nom");
        info.innerText = val.name + "\n" + "Chambre : " + val.n.toString();

        var icon = new Image(40, 40);
        icon.setAttribute("src", "static/image/chambreori.png");
        icon.setAttribute("style","right:0")
        icon.id = "icon-" + val.n.toString();

/*          mn par resident           ->
        var dur = document.createElement("div");
        dur.setAttribute("class", "contenu");
        dur.innerText = val.tmc.toString() + " mn";
        dur.id = "tmc-" + val.n.toString();*/


        box.appendChild(info);
        box.appendChild(icon);

        box.style.display = "flex";
        monitoring.appendChild(box);

    });
};




function update_css(data) {
    data.forEach(function creat(val) {
         if (val.lastEvent == "FALL") {

                evenement(val.name +" a chuté")

            } ;
        if (val.tmc<5) {

            $("#room-" + val.n.toString()).attr('data', val.tmc);
/*            $("#tmc-" + val.n.toString()).innerText = "0";*/
            $("#room-" + val.n.toString()).toggle(false);
            /*$("room-"+val.n.toString()).style.backgroundColor = "#f2f2f2"*/

        } else {

            $("#room-" + val.n.toString()).toggle(true);
            $("#room-" + val.n.toString()).attr("data", val.tmc);
            // $("#tmc-" + val.n.toString()).innerText = val.tmc;
            /*bleu :#1D7FB2; vert : #8C8910; rouge : #CA1725; gris :#f2f2f2;*/
            if (val.tmc == 30 ){
                            evenement( val.name + " est en activité depuis 30 mn");
                        }

            if (val.lastEvent == "BEDROOM") {
                    /* $("#room-" + val.n.toString() ).style.backgroundColor = "#f2f2f2";*/
                changeImage("icon-" + val.n.toString(), "static/image/chambreori.png");

            } else if (val.lastEvent == "BATHROOM") {


                /*$("#room-" + val.n.toString() ).style.backgroundColor = "#1D7FB2";*/
                changeImage("icon-" + val.n.toString(), "static/image/showerori.svg");

            }  else {

            };
        };
    });

/*    $(".floating-box").sort(sort_li).appendTo('#monitoring');
    function sort_li(a, b){
        return ($(b).data < ($(a).data ? 1 : -1;
    }*/

};

/*var myArray = $(".floating-box");
var count = 0;

// sort based on timestamp attribute
myArray.sort(function (a, b) {

    // convert to integers from strings
    a = parseInt($(a).attr("timestamp"), 10);
    b = parseInt($(b).attr("timestamp"), 10);
    count += 2;
    // compare
    if(a > b) {
        return 1;
    } else if(a < b) {
        return -1;
    } else {
        return 0;
    }
});

// put sorted results back on page
$("#results").append(myArray);
$("#calls").append(count+1);
*/
