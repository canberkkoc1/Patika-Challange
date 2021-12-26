import axios from "axios"

async function getData(id) {

    let urlUser = `https://jsonplaceholder.typicode.com/users/${id}`
    let urlPost = `https://jsonplaceholder.typicode.com/posts?userId=${id}`

    const userPromise = await axios.get(urlUser)
    const postPromise = await axios.get(urlPost)
    
    Promise.all([userPromise,postPromise]).then( res => {
        res.forEach(item => {
            console.log(item.data)
        })
    })

   
  
    

 

}


export default getData