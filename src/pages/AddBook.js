import Nav from "../components/home";
import { useState } from 'react'

function AddNewBook() {
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")

    async function addNewBooks() {         //we are defining the addNewBooks func 
                                           //here and we are calling it down in the button {addNewBooks}
        let item = { name, description }
        console.warn(name, description)
  
        

        let result = await fetch("http://localhost:8080/books/create", {
            method: "POST",
            body: JSON.stringify(item)
        })

        result = await result.json()
        alert("Your Book Has Been Added, Kindly CHeck it in the BookList Tab")


    }
    return (
        <div>
            <Nav />
            <div className="col-sm-6 offset-sm-3 h-1">
                <h1>Add Book</h1>
                <input type="text" value={name} onChange={(e) => setName(e.target.value)} className="form-control" placeholder="name" />
                <input type="text" value={description} onChange={(e) => setDescription(e.target.value)} className="form-control" placeholder="description" />
                <br></br>
                <button onClick={addNewBooks} className="w-100 btn btn-lg btn-primary">Add New Book</button>

            </div>
        </div>
    )
}

export default AddNewBook