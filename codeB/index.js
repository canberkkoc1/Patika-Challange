function SudokuQuadrantChecker(strArr) { 

    let newStr = [];
  
    strArr.forEach(function(item){
      newStr.push(item.replace(/[)(]/g,"").split(","))
  
    })
  
  
  
    // code goes here  
    console.log(newStr); 
  
  }


  SudokuQuadrantChecker([
    "(1,2,3,4,5,6,7,8,9)",
    "(x,x,x,x,x,x,x,x,x)",
    "(6,x,5,x,3,x,x,4,x)",
    "(2,x,1,1,x,x,x,x,x)",
    "(x,x,x,x,x,x,x,x,x)",
    "(x,x,x,x,x,x,x,x,x)",
    "(x,x,x,x,x,x,x,x,x)",
    "(x,x,x,x,x,x,x,x,x)",
    "(x,x,x,x,x,x,x,x,9)"
    ])