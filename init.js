db = db.getSiblingDB("messages");

db.createCollection("messages");

db.messages.insertMany([
    {
        content: "Hello!",
        recipient: "user1",
        sent_time: new Date().toISOString(),
        is_sent: true
    },
    {
        content: "Meeting at 3 PM.",
        recipient: "user2",
        sent_time: new Date().toISOString(),
        is_sent: true
    },
    {
        content: "Reminder: Project deadline.",
        recipient: "user3",
        sent_time: new Date().toISOString(),
        is_sent: true
    },
    {
        content: "Your order is confirmed.",
        recipient: "user4",
        sent_time: new Date().toISOString(),
        is_sent: true
    },
    {
        content: "Invoice attached.",
        recipient: "user5",
        sent_time: new Date().toISOString(),
        is_sent: true
    },
    {
        content: "Weekend plan?",
        recipient: "user6",
        sent_time: new Date().toISOString(),
        is_sent: true
    },
    {
        content: "Please call me.",
        recipient: "user7",
        sent_time: new Date().toISOString(),
        is_sent: false
    },
    {
        content: "Job interview update.",
        recipient: "user8",
        sent_time: new Date().toISOString(),
        is_sent: false
    },
    {
        content: "Password reset link.",
        recipient: "user9",
        sent_time: new Date().toISOString(),
        is_sent: false
    },
    {
        content: "Urgent: Please respond.",
        recipient: "user10",
        sent_time: new Date().toISOString(),
        is_sent: false
    }
]);