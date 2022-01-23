package models

import "gorm.io/gorm"

// const SECRET = process.env.SECRET

type role struct {
	user  string
	admin string
}

type User struct {
	gorm.Model
	username string
	password string
	loggedIn bool
	role
	// token: {
	//   type: DataTypes.VIRTUAL,
	//   get() {
	//     return jwt.sign({ username: this.username }, SECRET);
	//   },
	//   set(tokenObj) {
	//     let token = jwt.sign(tokenObj, SECRET);
	//     return token;
	//   },
	// },
	// capabilities: {
	//   type: DataTypes.VIRTUAL,
	//   get() {
	//     const acl = {
	//       user: ['read'],
	//       admin: ['read', 'create', 'update', 'delete'],
	//     };
	//     return acl[this.role];
	//   },
	// },

}

//  model.beforeCreate(async (user) => {
//    try {
//      const foundUser = await model.findOne({
//        where: { username: user.username },
//      });
//      if (foundUser) throw new Error('This username is taken');
//
//      let hashedPass = await bcrypt.hash(user.password, 10);
//      user.password = hashedPass;
//      user.loggedIn = true;
//    } catch (err) {
//      throw err;
//    }
//  });
//
//  model.authenticateBasic = async function (username, password) {
//    try {
//      const user = await this.findOne({ where: { username } });
//      if (!user) throw new Error('Please Register for an account first');
//      if (user.loggedIn) {
//        console.log(
//          `User ${user.username} is already logged in`,
//          user.loggedIn
//        );
//        throw new Error(`User ${user.username} is already logged in`);
//      } else {
//        const valid = await bcrypt.compare(password, user.password);
//        if (!valid) throw new Error('Invalid Username or Password');
//        await this.update(
//          { loggedIn: true },
//          {
//            where: {
//              username: username,
//            },
//          }
//        );
//      }
//
//      return user;
//    } catch (err) {
//      throw err;
//    }
//  };
//
//  model.authenticateToken = async function (token) {
//    try {
//      const parsedToken = jwt.verify(token, SECRET);
//      const user = this.findOne({ where: { username: parsedToken.username } });
//      if (!user) throw new Error('User Not Found');
//
//      return user;
//    } catch (err) {
//      console.log(err);
//      throw new Error(err.message);
//    }
//  };
//
//  model.logout = async function (username) {
//    try {
//      await this.update(
//        { loggedIn: false },
//        {
//          where: {
//            username: username,
//          },
//        }
//      );
//    } catch (err) {
//      console.log(err);
//    }
//
