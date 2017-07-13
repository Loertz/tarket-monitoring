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
        box.id = "room-" + val.RoomNumber.toString();
        box.setAttribute('data', val.Tmc);

        var info = document.createElement("div");
        info.setAttribute("class", "nom");
        info.innerText = val.Name + "\n" + "Chambre : " + val.RoomNumber.toString();

        var icon = new Image(40, 40);
        icon.setAttribute("src", "static/image/chambreori.png");
        icon.setAttribute("style","right:0")
        icon.id = "icon-" + val.RoomNumber.toString();

        box.appendChild(info);
        box.appendChild(icon);

        box.style.display = "flex";
        monitoring.appendChild(box);

    });
};

function update_css(val) {

    if (val.LastEvent == "FALL") {

        evenement(val.Name +" a chuté")

    } ;

    if (val.Tmc<5) {

        $("#room-" + val.RoomNumber.toString()).attr('data', val.Tmc);
        $("#room-" + val.RoomNumber.toString()).toggle(false);

    } else {

        $("#room-" + val.RoomNumber.toString()).toggle(true);
        $("#room-" + val.RoomNumber.toString()).attr("data", val.Tmc);
        // $("#Tmc-" + val.RoomNumber.toString()).innerText = val.Tmc;
        /*bleu :#1D7FB2; vert : #8C8910; rouge : #CA1725; gris :#f2f2f2;*/
        if (val.Tmc == 30 ){
                        evenement( val.Name + " est en activité depuis 30 mn");
                    }

        if (val.LastEvent == "BEDROOM") {
                /* $("#room-" + val.RoomNumber.toString() ).style.backgroundColor = "#f2f2f2";*/
            changeImage("icon-" + val.RoomNumber.toString(), "static/image/chambreori.png");
        if (val.BathRoomCount == 8 ){
                        evenement( val.Name + " est allé 8 fois en Salle de Bain");
                    }

        } else if (val.LastEvent == "BATHROOM") {


            /*$("#room-" + val.RoomNumber.toString() ).style.backgroundColor = "#1D7FB2";*/
            changeImage("icon-" + val.RoomNumber.toString(), "static/image/showerori.svg");

        }  else {

        };
    };
};
