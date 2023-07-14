import Navbar from "./components/navbar"
import getUser from "./utils/getUser"
import DailyQuestion from "./components/dailyQues";
import { useState } from "react";

export default function Home() {
  function logout() {
    togglelogin(false)
  }
  
  const [user, islogin] = getUser()
  const [isLoggedin, togglelogin] = useState<boolean>(islogin)
  
  let context
  if (isLoggedin) {
      context = (
          <div className="user-container">
              <div className="welcome-user">
                  <p>{user.username}</p>
              </div>
              <div className="daily-graph">
                  {/* // TODO: construct daily graph */} 
              </div>
              <div className="submission-history">
                  { /* // TODO: user submission history table */ }
              </div>
          </div>

      )
  } else {
        context = (
            <div className="login-container">
                < DailyQuestion />
            </div>
        )
    }

    return (
        <div className="home">
        <Navbar user={user} isLoggedin={isLoggedin} logout={logout}/>
        {context}
        </div>
    )
}