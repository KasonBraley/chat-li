const { users } = require('../models/index.js');

module.exports = function (io) {
  io.on('connection', (socket) => {
    console.log(`Socket ${socket.id} connected`);

    socket.on('message', ({ message, room }) => {
      try {
        io.to(room).emit('message', { username: socket.username, message });
      } catch (err) {
        console.log(err);
      }
    });

    socket.on('join', ({ room, username }) => {
      try {
        socket.join(room);
        if (username) {
          socket.username = username;
        }
        console.log(
          `${socket.username} (${socket.id}) just joined room [${room}]`
        );
        io.to(room).emit('user-connected', socket.username);
      } catch (err) {
        console.log(err);
      }
    });

    socket.on('leaveRoom', (room) => {
      try {
        socket.leave(room);
        console.log(`${socket.username} left room: ${room}`);
        io.to(room).emit('userLeftRoom', socket.username);
      } catch (err) {
        console.log(err);
      }
    });

    socket.on('quit', () => {
      try {
        console.log(`${socket.username} (${socket.id}) quit the chat`);
        socket.broadcast.emit('quit', socket.username);
        socket.disconnect(true);
      } catch (err) {
        console.log('Quit/logout error', err);
      }
    });

    socket.on('username', (username) => {
      try {
        socket.username = username;
      } catch (err) {
        console.log(err);
      }
    });

    socket.on('disconnecting', () => {
      try {
        users.logout(socket.username);
        if (socket.rooms.size === 1) {
          console.log(
            `${socket.username} (${socket.id}) quit the chat before joining a room`
          );
        }
      } catch (err) {
        console.log('Quit/logout error', err);
      }
    });

    socket.on('listUsers', () => {
      try {
        let allUsers = [];

        for (let [key, value] of io.of('/').sockets) {
          if (value.username) {
            allUsers.push(value.username);
          }
        }

        if (allUsers.length > 0) {
          io.to(socket.id).emit('listUsers', allUsers);
        }
      } catch (err) {
        console.log(err);
      }
    });

    socket.on('listRoomUsers', (room) => {
      try {
        let socketIds;
        let members = [];

        for (let [key, value] of io.of('/').adapter.rooms) {
          if (key.toLowerCase() === room) {
            socketIds = value;
          }
        }

        if (socketIds) {
          socketIds.forEach((socketId) => {
            let socketInRoom = io.of('/').sockets.get(socketId);
            members.push(socketInRoom.username);
          });

          io.to(socket.id).emit('listRoomUsers', { members, room });
        }
      } catch (err) {
        console.log(err);
      }
    });
  });
};
