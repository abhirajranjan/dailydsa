import {User} from "./models"

class user implements User {
    readonly username: string = ""
    readonly id: number = 0
    readonly profilePhoto: string = ""

    constructor(username: string, id: number, profilephoto: string) {
        this.username = username
        this.id = id
        this.profilePhoto = profilephoto
    }

    isValid(): boolean {
        return this.username !== ""  && this.id !== 0 ? true : false
    }
}

export function createUser(username: string, id: number, profilePhoto: string): User {
    return new user(username, id, profilePhoto)
}