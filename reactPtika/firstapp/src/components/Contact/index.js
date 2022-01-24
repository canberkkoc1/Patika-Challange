import {React,useState,useEffect} from 'react'
import List from './List'
import Form from './Form'

function COntacts() {

    const [contacts, setContacts] = useState([
        {
            fullName:"canberk",
        phoneNumber:50573639,
    },
        {
            fullName:"berce",
        phoneNumber:347843784373,
    },

    
    ])


    useEffect(() => {
        console.log(contacts)
    }, [contacts])


    return (
        <div>
            <h2>Contact</h2>
            <List contacts={contacts} />
            <Form addContacts = {setContacts} currentContacts = {contacts}/>
        </div>
    )
}

export default COntacts
