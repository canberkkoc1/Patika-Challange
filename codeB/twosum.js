


    function ParallelSums(arr) { 

        let newArr=[]

        let sum = 0
        
        for(let i = 0;i<arr.length;i++){
          for(let j = 1;j<arr.length;j++){
              console.log(arr[i],arr[j])
              sum = arr[i] + arr[j]
            if(sum === arr[i] + arr[j+1] ){
              newArr.push([arr[i],arr[j]])
            }
          }
        }
          // code goes here  
          console.log(newArr); 
        
        }




        ParallelSums([1,2,3,4])