import { TicketItemProps } from '../types';


const TicketItemDisplay = ({ ticket }: TicketItemProps) => {
    return (
        <>
            <td>{ticket.firstName} {ticket.lastName}</td>
            <td>{ticket.email}</td>
            <td>{ticket.issue}</td>
            <td style={{ width: '50px' }}>
                <span className={`priority-${ticket.priority?.toLowerCase()}`}>
                    {ticket.priority?.toLowerCase()}
                </span>
            </td>
        </>
    )
}

export default TicketItemDisplay;