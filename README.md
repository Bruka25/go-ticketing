# Go Ticketing Application

This is a simple ticket booking application written in Go. The application allows users to book tickets for a conference, validates user input, and sends a confirmation message asynchronously.

---

## Features

- **User Input Validation**: Ensures that user-provided data (name, email, ticket count) is valid.
- **Concurrency**: Sends ticket confirmation emails asynchronously using Goroutines.
- **Data Storage**: Stores booking information in memory using a slice of structs.
- **Real-Time Updates**: Tracks and updates the number of remaining tickets dynamically.

---

## Prerequisites

To run this application, you need:

- **Go**: Version 1.20 or higher installed on your system.
- A terminal or command prompt to execute the program.

---

## How to Run

1. Clone or download this repository to your local machine.
2. Navigate to the project directory:
   ```bash
   cd go-ticketing
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

---

## How It Works

1. **Greeting**:

   - The application welcomes users and displays the total and remaining tickets for the conference.

2. **User Input**:

   - Users are prompted to enter their first name, last name, email, and the number of tickets they want to book.

3. **Validation**:

   - The application validates the input:
     - First and last names must be at least 2 characters long.
     - Email must contain an `@` symbol.
     - The number of tickets must be greater than 0 and less than or equal to the remaining tickets.

4. **Booking**:

   - If the input is valid, the application:
     - Updates the remaining tickets.
     - Stores the booking information in memory.
     - Sends a confirmation email asynchronously.

5. **Completion**:
   - If all tickets are sold out, the application notifies users and stops accepting bookings.

---

## Code Overview

### Key Functions

- **`greetUsers()`**:
  Displays a welcome message and ticket information.

- **`getUserInput()`**:
  Collects user input for booking.

- **`validateInput()`**:
  Validates the user's input for correctness.

- **`bookTicket()`**:
  Updates the ticket count and stores booking information.

- **`sendTicket()`**:
  Simulates sending a ticket confirmation email asynchronously.

- **`getFirstName()`**:
  Extracts and displays the first names of all users who have booked tickets.

---

## Example Output

```plaintext
Welcome to Go conference the booking application!
We have a total of 50 tickets and 50 are still available
Get your tickets here to attend

Enter your first name:
John
Enter your last name:
Doe
Enter your Email:
john.doe@example.com
Enter number of tickets:
2

Thank you John Doe for booking 2 tickets. You will receive a confirmation email at john.doe@example.com
48 tickets remaining for Go conference
The first names of bookings are: [John]
```

---

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit a pull request.

---

## License

This project is open-source and available under the [MIT License](LICENSE).

---
