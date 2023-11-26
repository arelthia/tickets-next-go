'use client';
import { useState, ChangeEvent } from 'react';
import { Ticket } from '@/types';
import { useRouter } from 'next/navigation';



const AddTicket = () => {
    const router = useRouter();
    const [newTicket, setNewTicket] = useState<Ticket>({
        firstName: '',
        lastName: '',
        email: '',
        issue: '',
        priority: 'low'
    });
    const handleChange = (event: ChangeEvent<HTMLInputElement> | ChangeEvent<HTMLTextAreaElement> | ChangeEvent<HTMLSelectElement>) => {
        const { name, value } = event.target
        setNewTicket({
            ...newTicket,
            [name]: value
        })
    }

    const handleSubmit = async (event: ChangeEvent<HTMLFormElement>) => {
        event.preventDefault();
        const url = 'http://localhost:8000/api/v1/tickets';
        const response = await fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                firstName: newTicket.firstName,
                lastName: newTicket.lastName,
                email: newTicket.email,
                issue: newTicket.issue,
                priority: newTicket.priority
            })
        })
        const data: { created: number } = await response.json();
        if (data.created > 0) {
            handleBack();

        }

    }

    const handleBack = () => {

        const navigateTo = '/';
        router.push(navigateTo);
    }
    return (
        <div className='add-ticket-container'>
            <h2 className="page-title">Add a New Ticket</h2>
            <form onSubmit={handleSubmit} className='add-ticket'>

                <label>First Name
                    <input type="text" name="firstName" value={newTicket.firstName} onChange={handleChange} /></label>
                <label>Last Name
                    <input type="text" name="lastName" value={newTicket.lastName} onChange={handleChange} /></label>
                <label>Email</label>
                <input type="email" name="email" value={newTicket.email} onChange={handleChange} />
                <label>Issue</label>
                <textarea value={newTicket.issue} name="issue" onChange={handleChange}>{newTicket.issue}</textarea>
                <label>Priority</label>
                <span>
                    <select name="priority" id="" value={newTicket.priority} onChange={handleChange}>
                        <option value="low" >Low</option>
                        <option value="medium">Medium</option>
                        <option value="high">High</option>
                    </select>

                </span>



                <button type="submit">Submit</button>
                <button type="button" onClick={handleBack}>Cancel</button>

            </form>
        </div>
    )
}

export default AddTicket;