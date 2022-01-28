// let b = "canber - koc"
// let regex = /[^a-zA-Z0-9]/g
// console.log(b.replaceAll(regex,""))


function name(num) {
    if (num < 4) return num 
        

    console.log(name(num-1) + " ve " + name(num-2))

    return name(num-1) + name(num-2)


}


name(7)