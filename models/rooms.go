package models

import "gorm.io/gorm"

type Show struct {
	gorm.Model
	roomname string
	password string
}

// model.beforeCreate(async (room) => {
//   if (typeof myVar !== 'string' || !(myVar instanceof String))
//     room.password = room.password.toString();
//
//   let hashedPass = await bcrypt.hash(room.password, 10);
//   room.password = hashedPass;
// });
//
// model.authenticateBasic = async function (roomname, password) {
//   const room = await this.findOne({ where: { roomname } });
//
//   const valid = await bcrypt.compare(password, room.password);
//   if (valid) {
//     return room;
//   }
//   throw new Error('Invalid Room');
// };
