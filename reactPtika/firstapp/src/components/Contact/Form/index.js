import {React, useState} from 'react'

//! bu yontem daha efektif
// const initialFormValues =  {fullName:"",phoneNumber:""}

function Form({addContacts,currentContacts}) {

    const [form, setForm] = useState(
        {
            fullName:"",
            
            phoneNumber:"",
            

        }
    )

    const onChangeInput = (e) => {
        setForm({...form, [e.target.name] : e.target.value})
    }


    const onSubmit = (e) => {
        e.preventDefault()

        if(form.fullName ===  "" || form.phoneNumber === "" ){
            return false
        }

        //! eski kayıtlı contactları korumak için 
        addContacts([...currentContacts,form])

        console.log(form)

        setForm({fullName:"",phoneNumber:""})

        // setForm(initialFormValues)

        
    }

    return (
        <form onSubmit={onSubmit}>

            <div >
                <div>
                <input
                 type="text"  
                 name='fullName' 
                 placeholder='Fullname'  
                 value={form.fullName}
                 onChange={onChangeInput}/>

                </div>
                <div>
                <input 
                type="text" 
                name="phoneNumber" id=""  
                placeholder='Phone number' 
                value={form.phoneNumber}
                onChange={onChangeInput} 
                />
                

                </div>
                <button>ADD</button>
            </div>
        </form>
    )
}

export default Form
