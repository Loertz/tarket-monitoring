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


function update_css(val) {


    var box = document.createElement("div");
    box.setAttribute("class", "floating-box");
    box.id = "room-" + val.RoomNumber.toString();
    box.setAttribute('data', val.tmc);

    var info = document.createElement("div");
    info.setAttribute("class", "nom");
    info.innerText = val.name + "\n" + "Chambre : " + val.RoomNumber.toString();

    var icon = new Image(40, 40);
    icon.setAttribute("src", "static/image/chambreori.png");
    icon.setAttribute("style","right:0")
    icon.id = "icon-" + val.RoomNumber.toString();

    box.appendChild(info);
    box.appendChild(icon);

    box.style.display = "flex";
    monitoring.appendChild(box);

    if (val.lastEvent == "FALL") {

        evenement(val.name +" a chuté")

    } ;

    if (val.tmc<5) {

        $("#room-" + val.RoomNumber.toString()).attr('data', val.tmc);
        $("#room-" + val.RoomNumber.toString()).toggle(false);

    } else {

        $("#room-" + val.RoomNumber.toString()).toggle(true);
        $("#room-" + val.RoomNumber.toString()).attr("data", val.tmc);
        // $("#tmc-" + val.RoomNumber.toString()).innerText = val.tmc;
        /*bleu :#1D7FB2; vert : #8C8910; rouge : #CA1725; gris :#f2f2f2;*/
        if (val.tmc == 30 ){
                        evenement( val.name + " est en activité depuis 30 mn");
                    }

        if (val.lastEvent == "BEDROOM") {
                /* $("#room-" + val.RoomNumber.toString() ).style.backgroundColor = "#f2f2f2";*/
            changeImage("icon-" + val.RoomNumber.toString(), "static/image/chambreori.png");

        } else if (val.lastEvent == "BATHROOM") {


            /*$("#room-" + val.RoomNumber.toString() ).style.backgroundColor = "#1D7FB2";*/
            changeImage("icon-" + val.RoomNumber.toString(), "static/image/showerori.svg");

        }  else {

        };
    };
};
