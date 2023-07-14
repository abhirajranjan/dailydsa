import * as Popover from '@radix-ui/react-popover';
import {Cross2Icon } from '@radix-ui/react-icons';

import "./popover.css"

/**
 * popover displays a small popover section when trigger element is clicked.
 * @param trigger JSX Element that trigger popover to open
 * @param content JSX Element that will we shown in popover content area
 * @returns 
 */
export default function PopOver(trigger: JSX.Element, content: JSX.Element) {
    return <Popover.Root>
    <Popover.Trigger asChild>
      {trigger}
    </Popover.Trigger>
    <Popover.Portal>
      <Popover.Content className="PopoverContent" sideOffset={5}>
        {content}
        <Popover.Close className="PopoverClose" aria-label="Close">
          <Cross2Icon />
        </Popover.Close>
        <Popover.Arrow className="PopoverArrow" />
      </Popover.Content>
    </Popover.Portal>
  </Popover.Root>
}
