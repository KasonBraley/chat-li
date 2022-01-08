const base64 = require('base-64');
const { users, rooms } = require('../../models');

module.exports = async (req, res, next) => {
  try {
    if (req.body.isRoom) {
      let basic = req.headers.authorization.split(' ').pop();
      let [username, password] = base64.decode(basic).split(':');
      req.room = await rooms.authenticateBasic(username, password);
      next();
    } else {
      let basic = req.headers.authorization.split(' ').pop();
      let [user, pass] = base64.decode(basic).split(':');
      req.user = await users.authenticateBasic(user, pass);
      next();
    }
  } catch (err) {
    console.log(err);
    res.status(403).send(err.message);
  }
};
