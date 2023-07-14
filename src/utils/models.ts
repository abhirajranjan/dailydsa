export interface User {
    readonly username: string
    readonly id: number
    readonly profilePhoto: string

    // to check if user is valid or not 
    // must to used to check if login data exists or not
    isValid: () => boolean
}
