import { GoogleLogin } from '@react-oauth/google';

export default function GoogleLoginBtn() {
    return (
            <GoogleLogin
                onSuccess={credentialResponse => {
                    {/* // TODO: connect backend with login response */}
                    console.log(credentialResponse);
                }}
                onError={() => {
                    console.log('Login Failed');
                }}
            // useOneTap
            auto_select
            />
    )
}

