import {User} from "../utils/models"

import GoogleLoginBtn from "../components/signinCta";
import UserBadge from "./userBadge"

import "./navbar.css"

export default function Navbar({user, isLoggedin, logout}: {user: User | undefined, isLoggedin: boolean, logout: Function}) {
    let userProfile = <></> 
    if (user != undefined) {
        if (isLoggedin) {
            userProfile = <UserBadge user={user} logout={logout} />
        }
    }

    return (
        <nav className="navbar">
            <a href="/" className="title-container">
                <div className="title"> Daily DSA </div>
            </a>
            
            { isLoggedin && userProfile }

            <div className="logincta">
                { !isLoggedin &&  < GoogleLoginBtn />}
            </div>
        </nav>
    )
}



