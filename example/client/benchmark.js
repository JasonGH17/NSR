const net = require('net');

const { performance } = require('perf_hooks');

const s = new net.Socket();

const times = {
	createDB: [],
	add: [],
	get: [],
};

const create0 = performance.now();
s.connect({ port: 1766, host: 'localhost' }, () => {
	s.write('NSR');
	s.write('createDB "DB1" "schema1"');
});

s.on('data', (data) => {
	times.createDB.push(performance.now() - create0);
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

		const type = DB1commands[command].split(' ')[0];

		const cmd0 = performance.now();

		s.connect({ port: 1766, host: 'localhost' }, () => {
			s.write(database);
			s.write(DB1commands[command]);
		});

		s.on('data', (data) => {
			times[type].push(performance.now() - cmd0);

			data = data.toString('utf-8');
			console.log(`${data}\t Command: [${command}]`);
		});

		s.on('end', function () {
			console.log(`TCP Close\tCommand: [${command}]`);
		});
	}
});

s.on('end', function () {
	console.log('TCP Close\tDB: NSR');

});

setTimeout(DisplayTimes, 3000)
function DisplayTimes() {
	console.log()
	for (let arr in times) {
		console.log(`${arr} times:`);
		for (let time of times[arr]) console.log(`\t${time}ms`);
	}
}
