const net = require('net');

const commands = ['add test testdata ', 'get test '];

for (const command in commands) {
	const s = new net.Socket();

	s.connect({ port: 1766, host: 'localhost' }, () => {
		s.write(commands[command]);
	});

	s.on('data', (data) => {
		data = data.toString('utf-8');
		console.log(data !== 'Success' ? `${data}\t Command: [${command}]` : '');
	});

	s.on('end', function () {
		console.log(`TCP Close\tCommand: [${command}]`);
	});
}
