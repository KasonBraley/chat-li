module.exports = (capability) => {
  return (req, res, next) => {
    try {
      if (!req.user.capabilities.includes(capability)) {
        throw 'Access Denied';
      }
      next();
    } catch (e) {
      next(e);
    }
  };
};
