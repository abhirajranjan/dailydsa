import * as AlertDialog from '@radix-ui/react-alert-dialog';
import "./dialog.css"

/**
 * Generate a alert dialog that open in full screen with keyboard navigation.
 * Dialog box can be closed with esc key.
 * @param ctaBtnText text to show on trigger button that will trigger dialog
 * @param title title of the dialog box
 * @param description text to display inside dialog box
 * @param ctaCancelBtn JSX button to close the dialog 
 * @param ctaActionButton JSX button to do something on click 
 * @returns 
 */
export default function Dialog(ctaBtnText: string, title: string, description: string, ctaCancelBtn: JSX.Element, ...ctaActionButton: JSX.Element[]) {
    let ctaCancelJSX: JSX.Element // stores cancel button wrapped in AlertDialog.Cancel
    let ctaActionJSX: JSX.Element[] // stores action buttons wrapped in AlertDialog.Action

    ctaCancelJSX = (
      <AlertDialog.Cancel asChild>
        {ctaCancelBtn}
      </AlertDialog.Cancel>
    )

    ctaActionJSX = ctaActionButton.map<JSX.Element>((el: JSX.Element) => {
        return <AlertDialog.Action asChild>
            {el}
        </AlertDialog.Action>
    })

    return (
    <AlertDialog.Root>
    <AlertDialog.Trigger asChild>
      <button className="Button violet">{ctaBtnText}</button>
    </AlertDialog.Trigger>

    <AlertDialog.Portal>
      <AlertDialog.Overlay className="AlertDialogOverlay" />
      <AlertDialog.Content className="AlertDialogContent">
    
        <AlertDialog.Title className="AlertDialogTitle">{title}</AlertDialog.Title>
    
        <AlertDialog.Description className="AlertDialogDescription">
          {description}
        </AlertDialog.Description>
    
        <div style={{ display: 'flex', gap: 25, justifyContent: 'flex-end'}}>
            {ctaCancelJSX}
            {...ctaActionJSX}
        </div>
    
      </AlertDialog.Content>
    </AlertDialog.Portal>
  </AlertDialog.Root>
  )
}