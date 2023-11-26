import { useState, ChangeEvent } from 'react'
import { Ticket, TicketItemProps } from '../types';
import TicketItemDisplay from './TicketItemDisplay';




const TicketItem = ({ ticket }: TicketItemProps) => {
    const [editing, setEditing] = useState(false)
    const [editBtnText, setEditBtnText] = useState('Edit')
    const [newTicketData, setNewTicketData] = useState<Ticket>(ticket)

    const handleChange = (event: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLTextAreaElement> | ChangeEvent<HTMLSelectElement>) => {
        const { name, value } = event.target
        setNewTicketData({
            ...newTicketData,
            [name]: value
        })
    }

    const handleEditClick = async () => {
        if (editBtnText === 'Edit') {
            setEditing(true)
            setEditBtnText('Save')
        } else {
            const url = `http://localhost:8000/api/v1/tickets/${ticket.id}`
            const response = await fetch(url, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    firstName: newTicketData.firstName,
                    lastName: newTicketData.lastName,
                    email: newTicketData.email,
                    issue: newTicketData.issue,
                    priority: newTicketData.priority
                })
            })
            if (response.ok) {
                setEditing(false)
                setEditBtnText('Edit')
            }
            else {
                throw new Error('Failed to update ticket' + url)
            }
        }

    }

    const handleDeleteClick = async () => {
        const url = `http://localhost:8000/api/v1/tickets/${ticket.id}`
        const response = await fetch(url, {
            method: 'DELETE'
        })

    }

    return (
        <tr>
            {
                editing ? (
                    <>
                        <td>
                            <input style={{ width: '100px' }} type="text" name="firstName" value={newTicketData.firstName} onChange={handleChange} /> <input style={{ width: '100px' }} type="text" name="lastName" value={newTicketData.lastName} onChange={handleChange} />
                        </td>
                        <td><input type="email" name="email" value={newTicketData.email} onChange={handleChange} /></td>
                        <td><textarea name="issue" style={{ width: '250px' }} value={newTicketData.issue} onChange={handleChange}>{newTicketData.issue}</textarea></td>
                        <td style={{ width: '50px' }}>
                            <span className={`priority-${newTicketData.priority}`}>
                                <select name="priority" id="" value={newTicketData.priority} onChange={handleChange}>
                                    <option value="low" >Low</option>
                                    <option value="medium">Medium</option>
                                    <option value="high">High</option>
                                </select>

                            </span>
                        </td>
                    </>

                )

                    : <TicketItemDisplay ticket={ticket} />
            }


            <td>
                <span className="action-btn edit" onClick={handleEditClick}>{editBtnText}</span>
                {editing && <span className="action-btn cancel" onClick={() => { setEditing(false); setEditBtnText('Edit') }}>Cancel</span>}

                <span className="action-btn delete" onClick={handleDeleteClick}>Delete</span>
            </td>
        </tr>
    )
}

export default TicketItem