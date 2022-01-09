(function (){
    let userName = prompt("Lütfen isminizi giriniz")
    let headerName = document.querySelector("#headerTwo")

    let time = document.querySelector("#time")

    headerName.innerText = `${userName}! Hoşgeldin!`

    let days = {
        0: "pazar",
        1:"pazartes",
        2:"salı",
        3:"çarşamba",
        4:"perşembe",
        5:"cuma",
        6:"cumartesi",
    }
    
    setInterval(() => {
     let newDate = new Date()
         let day = newDate.getDay()
        let hour = newDate.getHours();
        let minutes = newDate.getMinutes();
        let seconds = newDate.getSeconds();


     
       time.innerText =  `${hour < 10 ? "0"+hour : hour} : ${minutes< 10 ? "0"+minutes : minutes} : ${seconds <10 ? "0"+seconds : seconds} ${days[day]} tarihinde kdoluyoruz front End development patikasına katıldınız`
    }, 1000);







})()