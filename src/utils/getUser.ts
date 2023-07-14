import {User} from "./models"
import { createUser } from "./user"

/**
 * 
 * @returns user object containing user information and boolean represent if user is logged-in or not
 */
export default function getUser(): readonly [User, boolean] {
    return [createUser("abhiraj ranjan", 1, "https://images.unsplash.com/photo-1511485977113-f34c92461ad9?ixlib=rb-1.2.1&w=128&h=128&dpr=2&q=80"), true] as const
}