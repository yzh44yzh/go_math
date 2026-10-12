# Define question feedback and retry behavior

Type: grilling
Status: resolved

## Question

How should each question transition through answer entry, feedback, and the next
question? Settle Enter/button behavior, malformed input, incorrect answers, whether
the correct answer is revealed before advancing, and repeated retry rounds.

Settle settings validation, allowable factor/count bounds, question-count capping,
progress and results wording, and abandoning or restarting a practice session.
Keep multiplication practice, the confirmed screen flow, and existing defaults.

## Answer

The user approved these interaction rules, with randomized retry order and
factor/count bounds adjusted as follows.

### Answer entry and feedback

- Enter or **Check answer** submits the current answer.
- Trim surrounding whitespace. Blank or non-integer input displays
  **Enter a whole number**; keep the question active without counting a mistake.
- A valid integer is checked once. Show **Correct** for a correct answer;
  otherwise show the correct multiplication result and record a mistake.
- Remain on feedback until Enter or **Next** advances. There is no automatic advance.
- After the final question's feedback, advancing opens the results screen.

### Questions and settings

- Both factors must be whole numbers from **1 through 9**, inclusive.
- The minimum factor must be less than or equal to the maximum factor.
- Question count must be a whole number from **1 through 100**, default **20**.
- Default factor range remains **1 through 9**.
- Generate all distinct ordered pairs in the selected inclusive range, shuffle
  them, and select up to the requested question count without replacement.
- **3 × 4** and **4 × 3** are separate questions.
- Cap the actual count at the available ordered pairs and explain the cap in the UI.
  Available count is `(maximum - minimum + 1)²`; the full 1–9 range has 81 questions.
  For example, requesting 100 questions from 1–9 starts a session of 81 questions.

### Progress, results, and retries

- Show progress such as **Question 3 of 20**, using the current round's actual count.
- Results show correct and incorrect counts for the completed round.
- Offer **Retry mistakes** only when the completed round has mistakes.
- Each retry round contains only mistakes from the immediately preceding round,
  shuffled into a random order. Shuffle anew for every retry round; a shuffle
  may coincidentally reproduce the previous order.
- Allow further retry rounds until all questions are answered correctly or the
  user chooses a new session or exits.

### Navigation

- **New session** returns to settings and retains the current settings values.
- During practice, **Back to settings** asks for confirmation before discarding
  progress. Cancelling that confirmation keeps the current practice state.
- Closing the window exits the application.

Concrete layout, focus, and the visual presentation of validation messages belong
to **Review the GUI screen layout**; these rules define behavior for that prototype.
