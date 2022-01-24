import React from 'react'
import { useState } from 'react'

function List({contacts}) {

    const [filterText,setFilterText] = useState("")


    const filtered = contacts.filter((item )=> {
       return Object.keys(item).some((key ) => {
           return  item[key]
           .toString()
           .toLowerCase()
           .includes(filterText.toLocaleLowerCase())
        }) 
    })


    console.log(filtered);


    return (
        <div>


        <input type="text" 
        value={filterText}
        onChange={(e) => setFilterText(e.target.value)}
        
        
        />


            <ul>

                {
                    filtered.map((contact,index) => {
                        return(
                            <li key={index}>{contact.fullName}</li>
                        )
                    })
                }
            </ul>
        </div>
    )
}

export default List
