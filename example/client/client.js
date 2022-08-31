const net = require('net');

const commands = ['createDB "DB1" "schema1"'];

const s = new net.Socket();

s.connect({ port: 1766, host: 'localhost' }, () => {
	s.write('NSR');
	s.write('createDB "DB1" "schema1"');
});

s.on('data', (data) => {
	data = data.toString('utf-8');
	console.log(data);

	const database = 'DB1';
	const DB1commands = [
		'add "schema1" "testsc1" "test1 schema1 data"',
		'add "schema2" "testsc2" "test2 schema2 data"',
		'get "schema2" "testsc2"',
		'get "schema1" "testsc1"',
	];

	for (const command in DB1commands) {
		const s = new net.Socket();

		s.connect({ port: 1766, host: 'localhost' }, () => {
			s.write(database);
			s.write(DB1commands[command]);
		});

		s.on('data', (data) => {
			data = data.toString('utf-8');
			console.log(`${data}\t Command: [${command}]`);
		});

		s.on('end', function () {
			console.log(`TCP Close\tCommand: [${command}]`);
		});
	}
});

s.on('end', function () {
	console.log('TCP Close');
});
