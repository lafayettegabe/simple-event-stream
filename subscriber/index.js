import { connect } from 'nats';

const natsUrl = 'nats://localhost:4222';
const subject = 'updates';

async function main() {
    const nc = await connect({ servers: natsUrl });
    console.log(`Connected to NATS at ${natsUrl}`);

    const sub = nc.subscribe(subject);
    console.log(`Subscribed to '${subject}'. Waiting for messages...`);

    (async () => {
        for await (const msg of sub) {
          console.log(`Received message on '${msg.subject}': ${msg.data}`);
        }
    })();
}

main()
