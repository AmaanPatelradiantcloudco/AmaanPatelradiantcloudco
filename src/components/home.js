import React from "react";
import { Button, NavDropdown } from "react-bootstrap";
import { Link } from "react-router-dom";

function Nav() {

    let user = JSON.parse(localStorage.getItem('user-info'))
    function logOut() {
        if (!localStorage.clear()) {
            return window.location.href = "/register"
        }
    }
    //to change a color in the class use text-warning,text-danger for different colours https://getbootstrap.com/docs/4.0/utilities/colors//
    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid ">
                <Link to="/" className="navbar-brand text-warning" >Home</Link> 
                <div>
                    <ul className="navbar-nav me-auto mb-2 mb-md-0 navbar_wrapper">
                        <div className="nav-item active">
                            {
                                localStorage.getItem('email', 'pwd') ?
                                    <>
                                        {/* <li className="nav-item active">
                                            <Link to="/" >Home</Link>
                                        </li> */}
                                        <li className="nav-item active text-danger">
                                       <Link to="/addNewbook" >Add New Book</Link>  
                                        </li>
                                        <li className="nav-item active text-warning">
                                            <Link to="/updatebook" >UpdateBook</Link>
                                        </li>
                                        <li className="nav-item active">
                                            <Link to="/bookslist" >BooksList</Link>
                                        </li>
                                        <li className="nav-item active">
                                            <Link to="/addNewauthor" >Add New Author</Link>
                                        </li>
                                        <li className="nav-item active">
                                            <Link to="/updateauthor" >UpdateAuthor</Link>
                                        </li>
                                        <li className="nav-item active">
                                            <Link to="/authorlist" >AuthorList</Link>
                                        </li>
                                    </>
                                    :
                                    <>


                                        <li className="nav-item active">
                                            <Link to="/login" >Login</Link>
                                        </li>
                                        <li className="nav-item active">
                                            <Link to="/register" >Register</Link>
                                        </li>

                                    </>
                            }
                    

                        </div >
                    </ul>
                </div>
            </div>
            {localStorage.getItem('email', 'pwd') ?

                <nav>
                    <NavDropdown title={user && user.name}>
                        <NavDropdown.Item onClick={logOut} >Logout</NavDropdown.Item>
                    </NavDropdown>
                </nav>
                : null
            }
        </nav>


    );
};
export default Nav;
