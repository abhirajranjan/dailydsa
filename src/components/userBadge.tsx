import {User} from "../utils/models"

import Avatar from "./base/avatar"
import Popover from "./base/popover"
import Dialog from "./base/dialog"

import "./userBadge.css"

export default function UserBadge({user, logout}: {
    user: User,
    logout: Function
    }) {
        let description // contain all the description of the content. eg avatar and name
        let actions // contains action buttons

        let avatarBtn = <button> {Avatar({Src: user.profilePhoto, ClassName: ""})} </button> // avatar to trigger popover
        let popavatar = Avatar({Src: user.profilePhoto, ClassName: "popavatar"}) // avatar in popover
        
        let logoutDialogCancelBtn = <button className="Button mauve">Cancel</button>
        let logoutDialoglogoutBtn = <button className="Button red" onClick={() => logout()}>logout</button>
        let logoutDialog  = Dialog("logout", "logout", "Are you sure that you want to logout ?", logoutDialogCancelBtn, logoutDialoglogoutBtn)

        description = <div className="description"> 
            {popavatar}
            <p className="username">{user.username}</p>
        </div>

        actions = <div className="actions">
            {logoutDialog}
        </div>

        let content = (
            <div style={{"display": "flex", "flexDirection": "column", "gap": "30px"}}>
                {description}
                {actions}
            </div>
        )

        return Popover(avatarBtn, content)
}
