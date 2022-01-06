

function hoursDec(str){


    let [firstStr, secondStr] = str.split("-")

    function convert(h){
        
        let timePer = h.slice(-2).toLowerCase()

        let hours = h.replace(/[a-z]./g,"")

        let [hour,minute] = hours.split(":")

        if(hour === "12"){
            hour = "00"
        }

        if(timePer === "pm"){
            hour = parseInt(hour,10) + 12
        }

        return `${hour}:${minute}`



    }

    let first = convert(firstStr)
    let second = convert(secondStr)

    let [firstHour,firstMin] = first.split(":")
    let [secondHour,secondMin] = second.split(":")
    
    let date1 = new Date(1994,0,1,+firstHour,+firstMin)
    let date2 = new Date(1994,0,1,+secondHour,+secondMin)

    let finalResult = (date1-date2)/60000
    
    console.log(finalResult)
    
    
}


hoursDec("1:20pm-11:00am")
