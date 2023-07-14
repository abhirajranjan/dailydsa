import * as avatar from '@radix-ui/react-avatar';
import "./avatar.css"


/**overriding class ClassName should have !important to work because of 
css precedance in inline classnames
*/

/**
 * give the avatar with some predefined css.
 * 
 * use !important in ClassName css to override default 
 * @param src - img source string of image that will be displayed
 * @param className - name of class that will be applied to the avatar. 
 * @returns {JSX.Element}
 */
export default function Avatar({Src, ClassName} : {Src: string, ClassName: string}) {
    return (
        <avatar.Root className={`${ClassName} AvatarRoot`}>
            <avatar.Image
            className="AvatarImage"
                src={Src}
                alt="Pedro Duarte"
            />
            <avatar.Fallback className="AvatarFallback" delayMs={600}>
                JD
            </avatar.Fallback>
        </avatar.Root>
    )
}