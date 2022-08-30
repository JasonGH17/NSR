const net = require('net');

const commands = ['get "schema1" "test"', 'get "schema2" "testsc2"', 'get "schema2" "test2"'];

for (const command in commands) {
	const s = new net.Socket();

	s.connect({ port: 1766, host: 'localhost' }, () => {
		s.write(commands[command]);
	});

	s.on('data', (data) => {
		data = data.toString('utf-8');
		console.log(`${data}\t Command: [${command}]`);
	});

	s.on('end', function () {
		console.log(`TCP Close\tCommand: [${command}]`);
	});
}
